package pur

// SupplierAsnInfoVO 
type SupplierAsnInfoVO struct {
    // 关闭原因
    CloseReason   string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
    // asn状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // asn备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 实际签收人
    ReceivedBy   string `json:"received_by,omitempty" xml:"received_by,omitempty"`
    // 签收人
    Receipient   string `json:"receipient,omitempty" xml:"receipient,omitempty"`
    // 国际贸易交付地点
    ConsigneeAddress   string `json:"consignee_address,omitempty" xml:"consignee_address,omitempty"`
    // 国内交付地点
    DeliveryAddress   string `json:"delivery_address,omitempty" xml:"delivery_address,omitempty"`
    // 签收方
    Consignee   string `json:"consignee,omitempty" xml:"consignee,omitempty"`
    // 木箱号
    Cass   string `json:"cass,omitempty" xml:"cass,omitempty"`
    // 集装箱号
    ContainerNo   string `json:"container_no,omitempty" xml:"container_no,omitempty"`
    // 装箱单号
    PackingListNo   string `json:"packing_list_no,omitempty" xml:"packing_list_no,omitempty"`
    // 国际贸易术语
    Incoterms   string `json:"incoterms,omitempty" xml:"incoterms,omitempty"`
    // 物流单号
    TrackingNum   string `json:"tracking_num,omitempty" xml:"tracking_num,omitempty"`
    // 承运方联系电话
    CarrierContactPhone   string `json:"carrier_contact_phone,omitempty" xml:"carrier_contact_phone,omitempty"`
    // 承运方联系邮箱
    CarrierContactEmail   string `json:"carrier_contact_email,omitempty" xml:"carrier_contact_email,omitempty"`
    // 承运方联系人
    CarrierContact   string `json:"carrier_contact,omitempty" xml:"carrier_contact,omitempty"`
    // 承运人
    Carrier   string `json:"carrier,omitempty" xml:"carrier,omitempty"`
    // 配送方式
    ShipmentType   string `json:"shipment_type,omitempty" xml:"shipment_type,omitempty"`
    // 实际发货日期
    ActualArrivalDate   string `json:"actual_arrival_date,omitempty" xml:"actual_arrival_date,omitempty"`
    // 预计到货日期
    EstimatedArrivalDate   string `json:"estimated_arrival_date,omitempty" xml:"estimated_arrival_date,omitempty"`
    // 采购组织代码
    DemanderPurchaseOrgCode   string `json:"demander_purchase_org_code,omitempty" xml:"demander_purchase_org_code,omitempty"`
    // Ou代码
    OuCode   string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
    // 供应商名称
    SupplierName   string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
    // 供应商编码
    SupplierCode   string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
    // asn行信息
    AsnItemList   []SupplierAsnItem `json:"asn_item_list,omitempty" xml:"asn_item_list>supplier_asn_item,omitempty"`
}
