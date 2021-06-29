package security

// RpAuditMaterialDetail 
type RpAuditMaterialDetail struct {
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // display
    Display   string `json:"display,omitempty" xml:"display,omitempty"`
    // intercept
    Intercept   bool `json:"intercept,omitempty" xml:"intercept,omitempty"`
    // materialType
    MaterialType   string `json:"material_type,omitempty" xml:"material_type,omitempty"`
    // security
    Security   bool `json:"security,omitempty" xml:"security,omitempty"`
    // suggestion
    Suggestion   string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
    // text
    Text   string `json:"text,omitempty" xml:"text,omitempty"`
    // type
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
}
