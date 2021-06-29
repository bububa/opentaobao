package mos

// CallDispatcherRespDO 
type CallDispatcherRespDO struct {
    // shortId
    ShortId   string `json:"short_id,omitempty" xml:"short_id,omitempty"`
    // gotCode
    GotCode   string `json:"got_code,omitempty" xml:"got_code,omitempty"`
    // packageCode
    PackageCode   string `json:"package_code,omitempty" xml:"package_code,omitempty"`
    // fulfillPlanId
    FulfillPlanId   int64 `json:"fulfill_plan_id,omitempty" xml:"fulfill_plan_id,omitempty"`
}
