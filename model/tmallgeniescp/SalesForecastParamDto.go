package tmallgeniescp

// SalesForecastParamDto 
type SalesForecastParamDto struct {
    // 扩展参数
    ExtendJson   string `json:"extend_json,omitempty" xml:"extend_json,omitempty"`
    // 租户
    Tenant   string `json:"tenant,omitempty" xml:"tenant,omitempty"`
    // 日期
    KeyFigureDate   string `json:"key_figure_date,omitempty" xml:"key_figure_date,omitempty"`
    // 渠道
    CustId   string `json:"cust_id,omitempty" xml:"cust_id,omitempty"`
    // 物料号
    PrdId   string `json:"prd_id,omitempty" xml:"prd_id,omitempty"`
    // 产管提报预测数量-月度
    SalesForecastQtyPmPerMonth   string `json:"sales_forecast_qty_pm_per_month,omitempty" xml:"sales_forecast_qty_pm_per_month,omitempty"`
    // 产管-提报预测数量
    SalesForecastQtyPm   string `json:"sales_forecast_qty_pm,omitempty" xml:"sales_forecast_qty_pm,omitempty"`
    // 销管提报预测数量
    SalesForecastQtySl   string `json:"sales_forecast_qty_sl,omitempty" xml:"sales_forecast_qty_sl,omitempty"`
}
