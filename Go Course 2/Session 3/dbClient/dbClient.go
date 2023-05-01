package dbClient

import "github.com/stretchr/testify/mock"

type DBClient interface{
	Query(sql string) any
}


//concrete implementation
type MySQLDBClient struct{

}

func (mysql MySQLDBClient) Query(sql string) any{
	return ""
}

//stub
type MockDBClient struct{
	mock.Mock
}

func (dbMock MockDBClient) Query(sql string) any{
	args := dbMock.Called(sql)
	return args.Get(0).(any)
}