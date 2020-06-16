package models

func AllDataGV() ([]*Test, error) {
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tst := make([]*Test, 0)
	for rows.Next() {
		ts := new(Test)
		err := rows.Scan(&ts.Id, &ts.Title)
		if err != nil {
			return nil, err
		}
		tst = append(tst, ts)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tst, nil
}
