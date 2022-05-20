package storage

import (
	"log"

	"github.com/ferdunosmanov/inspectors/pkg/adding"
	"github.com/gocql/gocql"
)

type Storage struct {
	db *gocql.Session
}

func SetupStorage() (*Storage, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	cluster.Keyspace = "inspection"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{db: session}, nil
}

func (s *Storage) GetAllProductNames() ([]adding.Product, error) {
	var product adding.Product
	var products []adding.Product
	iter := s.db.Query(`SELECT tapID, timestamp, date, latitude, longitude, staffid, staffname, shiftid, deviceid, vehicleid, metro, linetype, coursenumber, mediatype, mediaid, medianumber, status, line FROM inspections`).Iter()
	for iter.Scan(
		&product.TapId,
		&product.Timestamp,
		&product.Date,
		&product.Latitude,
		&product.Longitude,
		&product.StaffId,
		&product.StaffName,
		&product.ShiftId,
		&product.DeviceId,
		&product.VehicleId,
		&product.Metro,
		&product.LineType,
		&product.CourseNumber,
		&product.MediaType,
		&product.MediaId,
		&product.MediaNumber,
		&product.Status,
		&product.Line,
	) {
		products = append(products, product)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return products, nil
}

func (s *Storage) GetAllRegistrations() ([]adding.Registrations, error) {
	var product adding.Registrations
	var products []adding.Registrations
	iter := s.db.Query(`SELECT date, event_id, datetime, staff_id, shift_id, longitude, latitude, vehicle_id, line_id, route_id, course_no, stop_id, manual_sync, cluster FROM registrations`).Iter()
	for iter.Scan(
		&product.Date,
		&product.Event_id,
		&product.Datetime,
		&product.Staff_id,
		&product.Shift_id,
		&product.Longitude,
		&product.Latitude,
		&product.Vehicle_id,
		&product.Line_id,
		&product.Route_id,
		&product.Course_no,
		&product.Stop_id,
		&product.Manual_sync,
		&product.Cluster,
	) {
		products = append(products, product)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return products, nil
}
