package adding

type Product struct {
	TapId        string  `json:"tapid"`
	Timestamp    int64   `json:"timestamp"`
	Date         string  `json:"date"`
	StaffId      int     `json:"staffid"`
	StaffName    string  `json:"staffname"`
	ShiftId      int     `json:"shiftid"`
	DeviceId     string  `json:"deviceid"`
	VehicleId    string  `json:"vehicleid"`
	Metro        bool    `json:"metro"`
	LineType     string  `json:"linetype"`
	CourseNumber int     `json:"coursenumber"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	MediaType    string  `json:"mediatype"`
	MediaId      string  `json:"mediaid"`
	MediaNumber  string  `json:"medianumber"`
	Status       string  `json:"status"`
	Line         string  `json:"line"`
}

type Registrations struct {
	Date        string  `json:"date"`
	Event_id    string  `json:"event_id"`
	Datetime    int64   `json:"datetime"`
	Staff_id    int     `json:"staff_id"`
	Shift_id    int     `json:"shift_id"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Vehicle_id  string  `json:"vehicle_id"`
	Line_id     string  `json:"line_id"`
	Route_id    string  `json:"route_id"`
	Course_no   int     `json:"course_no"`
	Stop_id     string  `json:"stop_id"`
	Manual_sync bool    `json:"manual_sync"`
	Cluster     int     `json:"cluster"`
}
