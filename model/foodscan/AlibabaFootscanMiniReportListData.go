package foodscan

// AlibabaFootscanMiniReportListData 
type AlibabaFootscanMiniReportListData struct {
    // 页码
    PageNum   int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
    // 页长
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 列表数据
    List   []FeetModelMobileReportVo `json:"list,omitempty" xml:"list>feet_model_mobile_report_vo,omitempty"`
}
