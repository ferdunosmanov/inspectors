package adding

type Product struct {
	MediaID            string `json:"mediaId",omitempty`
	DocumentID         int64  `json:"documentId"`
	ProductType        string `json:"productType"`
	ProductName        string `json:"productName"`
	ProductValidFrom   string `json:"productValidFrom"`
	ProductValidTo     string `json:"productValidTo"`
	ProductVehicleType string `json:"productVehicleType"`
	ProductLines       string `json:"productLines"`
	Metro              bool   `json:"metro"`
	IsCancelled        bool   `json:"isCancelled"`
}
