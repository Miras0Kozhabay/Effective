package repository

import (
	"database/sql"
	"fmt"

	"subscriptions-service/internal/model"

	"github.com/google/uuid"
)

type SubscriptionRepo struct {
	DB *sql.DB
}

func NewSubscriptionRepo(db *sql.DB) *SubscriptionRepo {
	return &SubscriptionRepo{DB: db}
}

func (r *SubscriptionRepo) Create(s model.Subscription) error {
	_, err := r.DB.Exec(
		"INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date) VALUES ($1,$2,$3,$4,$5,$6)",
		s.ID, s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate,
	)
	return err
}

func (r *SubscriptionRepo) GetAll() ([]model.Subscription, error) {
	rows, err := r.DB.Query("SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []model.Subscription
	for rows.Next() {
		var s model.Subscription
		err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
		if err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func (r *SubscriptionRepo) GetByID(id string) (*model.Subscription, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var s model.Subscription
	err = r.DB.QueryRow("SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id=$1", u).
		Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SubscriptionRepo) Update(id string, s model.Subscription) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec("UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5 WHERE id=$6",
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, u)
	return err
}

func (r *SubscriptionRepo) Delete(id string) error {
	u, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec("DELETE FROM subscriptions WHERE id=$1", u)
	return err
}

func (r *SubscriptionRepo) GetTotal(userID, serviceName, start, end string) (int, error) {
	query := "SELECT COALESCE(SUM(price),0) FROM subscriptions WHERE 1=1"
	args := []interface{}{}
	i := 1

	if userID != "" {
		query += fmt.Sprintf(" AND user_id=$%d", i)
		args = append(args, userID)
		i++
	}
	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name=$%d", i)
		args = append(args, serviceName)
		i++
	}
	if start != "" {
		query += fmt.Sprintf(" AND start_date >= $%d", i)
		args = append(args, start)
		i++
	}
	if end != "" {
		query += fmt.Sprintf(" AND start_date <= $%d", i)
		args = append(args, end)
		i++
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
