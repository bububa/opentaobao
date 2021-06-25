package qimen

// Package 
type Package struct {

    // 物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通 (ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小 包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
    LogisticsCode   string `json:"logisticsCode,omitempty"`

    // 物流公司名称
    LogisticsName   string `json:"logisticsName,omitempty"`

    // 运单号
    ExpressCode   string `json:"expressCode,omitempty"`

    // 包裹编号
    PackageCode   string `json:"packageCode,omitempty"`

    // 包裹长度(单位：厘米)
    Length   string `json:"length,omitempty"`

    // 包裹宽度(单位：厘米)
    Width   string `json:"width,omitempty"`

    // 包裹高度(单位：厘米)
    Height   string `json:"height,omitempty"`

    // 包裹理论重量(单位：千克)
    TheoreticalWeight   string `json:"theoreticalWeight,omitempty"`

    // 包裹重量(单位：千克)
    Weight   string `json:"weight,omitempty"`

    // 包裹体积(单位：升)
    Volume   string `json:"volume,omitempty"`

    // 发票号
    InvoiceNo   string `json:"invoiceNo,omitempty"`

    // 包材信息
    PackageMaterialList   []PackageMaterial `json:"packageMaterialList,omitempty"`

    // 商品列表
    Items   []Item `json:"items,omitempty"`

    // 签收人姓名, string (50) ，必填
    SignUserName   string `json:"signUserName,omitempty"`

    // 当前状态操作时间, string (19) , YYYY-MM-DD HH:MM:SS
    SignTime   string `json:"signTime,omitempty"`

    // 状态, sign-已签收string (50)
    Status   string `json:"status,omitempty"`

    // 备注, string (500) ,
    Remarks   string `json:"remarks,omitempty"`

}
