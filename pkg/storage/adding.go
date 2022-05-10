package storage

import "log"

func (s *Storage) AddProduct(c adding.Candy) (string, error) {
	id := uuid.New().String()
	if err := s.db.Query(`INSERT INTO products (MediaID, DocumentID, ProductType, ProductName, ProductValidFrom, ProductValidTo, ProductVehicleType, ProductLines, Metro, IsCancelled) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		id, c.MediaID, c.DocumentID, c.ProductType, c.ProductName, c.ProductValidFrom, c.ProductValidTo, c.ProductVehicleType, c.ProductLines, c.Metro, c.IsCancelled).Exec(); err != nil {
		log.Println("Error while trying to save to DB: ", err)
		return "", err
	}
}
