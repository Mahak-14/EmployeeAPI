package migrations

const (
	createEmployeeTable = "CREATE TABLE IF NOT EXISTS employee (ID INT, NAME VARCHAR(255), DEPT VARCHAR(255));"
	dropEmployeeTable   = "DROP TABLE IF EXISTS employee"
)
