package models

import (
	"log"
	"time"

	"github.com/ClementTeyssa/3PJT-API/config"
)

type ToReward struct {
	ID        int       `json:"id" validate:"omitempty,uuid"`
	Adress    string    `json:"adress"`
	Number    int       `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ToRewards []ToReward

func NewToReward(toReward *ToReward) {
	if toReward == nil {
		log.Panic(toReward)
	}
	toReward.CreatedAt = time.Now()
	toReward.UpdatedAt = time.Now()
	err := config.GetDb().QueryRow("INSERT INTO torewards (adress, number, created_at, updated_at) VALUES ($1,$2,$3,$4) RETURNING id;", toReward.Adress, toReward.Number, toReward.CreatedAt, toReward.UpdatedAt).Scan(&toReward.ID)

	if err != nil {
		log.Panic(err)
	}
}

func FindToRewardByAdress(adress string) *ToReward {
	var toReward ToReward
	row := config.GetDb().QueryRow("SELECT * FROM torewards WHERE adress = $1;", adress)
	err := row.Scan(&toReward.ID, &toReward.Adress, &toReward.Number, &toReward.CreatedAt, &toReward.UpdatedAt)

	if err != nil {
		return nil
	}

	return &toReward
}

func CountToRewardByAdress(adress string) int {
	rows, err := config.GetDb().Query("SELECT COUNT(*) as count FROM torewards WHERE adress = $1;", adress)

	if err != nil {
		log.Panic(err)
	}

	count := 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Panic(err)
		}
	}

	return count
}

func AllToReward() *ToRewards {
	var toRewards ToRewards
	rows, err := config.GetDb().Query("SELECT * FROM torewards")
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var toReward ToReward
		err := rows.Scan(&toReward.ID, &toReward.Adress, &toReward.Number, &toReward.CreatedAt, &toReward.UpdatedAt)
		if err != nil {
			log.Panic(err)
		}
		toRewards = append(toRewards, toReward)
	}
	return &toRewards
}

func UpdateToReward(toReward *ToReward) {
	toReward.UpdatedAt = time.Now()
	stmt, err := config.GetDb().Prepare("UPDATE torewards SET adress=$1, number=$2, aupdated_at=$3 WHERE id=$4;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(toReward.Adress, toReward.Number, toReward.UpdatedAt)
	if err != nil {
		log.Panic(err)
	}
}
