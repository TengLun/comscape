package comscape

func MakeDAO(debug bool) (DataAccessor, error) {
	if debug == true {
		var da DAOFake
		return da, nil
	}
	var da DAO
	return da, nil
}
