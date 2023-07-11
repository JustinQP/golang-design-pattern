package abstractfactory

import "fmt"

/*
1、抽象工厂（Abstract Factory）：定义了一组用于创建相关产品对象的方法，每个方法对应一个产品对象的创建。
2、具体工厂（Concrete Factory）：实现了抽象工厂接口，负责创建一套相关产品对象。
3、抽象产品（Abstract Product）：定义了一个产品对象的接口，具体的产品类需要实现这个接口。
4、具体产品（Concrete Product）：实现了抽象产品接口，是抽象工厂模式创建的对象。
*/

// OrderMainDAO 为订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情纪录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

////////////////// RDB /////////////////////////////
// RDBMainDAO 关系型数据库的OrderMainDAO实现
type RDBMainDAO struct{}

// SaveOrderMain ...
func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct{}

// SaveOrderDetail ...
func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

////////////////// XML /////////////////////////////
// XMLMainDAO XML存储
type XMLMainDAO struct{}

// SaveOrderMain ...
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XMLDetailDAO XML存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

// XMLDAOFactory 是XML 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
