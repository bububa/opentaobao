package tmallcampus

// CollegeAuditStatusResponse 
type CollegeAuditStatusResponse struct {

    // 认证状态描述;一共四种状态audit_name和audit_code对应关系如下：1.认证是学生 audit_pass;2.认证非学生 audit_no_pass;3.审核中 audit_wait;4.未认证 audit_no_record;
    
    AuditName   string `json:"audit_name,omitempty" xml:"audit_name,omitempty"`
    

    // 认证状态编码
    
    AuditCode   string `json:"audit_code,omitempty" xml:"audit_code,omitempty"`
    

}
