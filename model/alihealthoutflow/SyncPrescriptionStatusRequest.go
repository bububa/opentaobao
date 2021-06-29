package alihealthoutflow

// SyncPrescriptionStatusRequest 
type SyncPrescriptionStatusRequest struct {
    // 阿里将康处方号(非空)
    RxNo   string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
    // 处方状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
}
