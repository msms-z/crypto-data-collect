package model

type ContractData struct {
	ContractDecimal int    `json:"contract_decimal" gorm:"column:contract_decimal;default:18;comment:合约精度"`
	ContractAddress string `json:"contract_address" gorm:"column:contract_address;unique;primary_key;comment:合约地址"`
	ContractAbiData string `json:"contract_abi_data" gorm:"column:contract_abi_data;comment:ABI"`
}

func (ContractData) TableName() string {
	return "contract_data"
}
