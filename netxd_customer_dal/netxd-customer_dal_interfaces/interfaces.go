package netxdcustomerdalinterfaces

import netxdcustomerdalmodels "github.com/buvanesh2002/netxd_dal/netxd_customer_dal/netxd_customer_dal_models"



type ICustomer interface {
	CreateCustomer(user *netxdcustomerdalmodels.Customer) (*netxdcustomerdalmodels.CustomerResponse, error)
}
