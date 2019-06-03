package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ClementTeyssa/3PJT-API3/models"

	"github.com/ClementTeyssa/3PJT-API3/helper"
)

type Node struct {
	Adress string `json:"adress"`
}

type ToAddNodes []Node

type TransactionsSend struct {
	To []string `json:"transactions"`
}

func AddReward(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "ioutil.ReadAll(r.Body)")
		return
	}

	var nodes ToAddNodes
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		helper.ErrorHandlerHttpRespond(w, "json.Unmarshal(body, &nodes)")
		return
	}

	var transactionsSend *TransactionsSend
	for _, node := range nodes {
		if models.CountToRewardByAdress(node.Adress) == 0 {
			toReward := models.FindToRewardByAdress(node.Adress)
			if toReward.Number == 4 {
				toReward.Number = 0
				doReward(node.Adress)
				transactionsSend.To = append(transactionsSend.To, toReward.Adress)
			} else {
				toReward.Number = toReward.Number + 1
			}
			models.UpdateToReward(toReward)
		} else {
			var toReward *models.ToReward
			toReward.Adress = node.Adress
			toReward.Number = 1
			models.NewToReward(toReward)
		}
	}

	json.NewEncoder(w).Encode(transactionsSend)
}

func doReward(adress string) {
	// lance une demande de transaction à l'api2
}
