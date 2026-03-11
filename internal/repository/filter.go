package repository

import (
	"fmt"
	"subscriptions-service/internal/model"
)

func (r *SubscriptionRepo) GetAllFiltered(userID, serviceName string, minPrice, maxPrice int, sort string) ([]model.Subscription, error) {
	query := "SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	argID := 1

	if userID != "" {
		query += fmt.Sprintf(" AND user_id = $%d", argID)
		args = append(args, userID)
		argID++
	}
	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name ILIKE $%d", argID)
		args = append(args, "%"+serviceName+"%")
		argID++
	}
	if minPrice > 0 {
		query += fmt.Sprintf(" AND price >= $%d", argID)
		args = append(args, minPrice)
		argID++
	}
	if maxPrice > 0 {
		query += fmt.Sprintf(" AND price <= $%d", argID)
		args = append(args, maxPrice)
		argID++
	}

	if sort == "price_asc" {
		query += " ORDER BY price ASC"
	} else if sort == "price_desc" {
		query += " ORDER BY price DESC"
	} else if sort == "start_date_asc" {
		query += " ORDER BY start_date ASC"
	} else if sort == "start_date_desc" {
		query += " ORDER BY start_date DESC"
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subs := []model.Subscription{}
	for rows.Next() {
		var s model.Subscription
		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}
