package admarket

// Network 
type Network struct {
    // 基站id
    CellularId   string `json:"cellular_id,omitempty" xml:"cellular_id,omitempty"`
    // 运营商(UNKNOWN_OPERATOR/CHINA_MOBILE/CHINA_TELECOM/CHINA_UNICOM/OTHER_OPERATOR)
    OperatorType   string `json:"operator_type,omitempty" xml:"operator_type,omitempty"`
    // 网络类型(CONNECTION_UNKNOWN/CELL_UNKNOWN/CELL_2G/CELL_3G/CELL_4G/CELL_5G/WIFI/ETHERNET/NEW_TYPE)
    ConnectionType   string `json:"connection_type,omitempty" xml:"connection_type,omitempty"`
    // ip地址
    Ip   string `json:"ip,omitempty" xml:"ip,omitempty"`
}
