package store

const (
	createQuery  = `INSERT INTO employee (id,name,dept)values(?,?,?)`
	updateQuery  = `UPDATE employee SET name=?, dept=? WHERE id=?`
	getByIDQuery = `SELECT id,name,dept FROM employee WHERE id=?`
	deleteQuery  = `Delete from employee where id =?`
)
