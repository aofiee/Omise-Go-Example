package models

// OmiseKey Struct
type OmiseKey struct {
	ID        int
	PublicKey string
	SecretKey string
}

// func FindEmployee(number int) (bool, *Employee) {
// 	obj, err := dbmap.Get(Employee{}, number)
// 	emp := obj.(*Employee)

// 	if err != nil {
// 		log.Print("ERROR findEmployee: ")
// 		log.Println(err)
// 	}

// 	return (err == nil), emp
// }
