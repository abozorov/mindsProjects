package models

type Cell struct {
	ID         int    `json:"id"`
	Zone       string `json:"zone"`
	Row        int    `json:"row"`
	AdressCode string `json:"adress_code" db:"adress_code"`
}
