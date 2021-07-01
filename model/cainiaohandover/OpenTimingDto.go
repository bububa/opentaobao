package cainiaohandover

// OpenTimingDto 
type OpenTimingDto struct {
    // 展示文案
    DisplayText   string `json:"display_text,omitempty" xml:"display_text,omitempty"`
    // 最慢时效
    SlowestTiming   int64 `json:"slowest_timing,omitempty" xml:"slowest_timing,omitempty"`
    // 最快时效
    FastTiming   int64 `json:"fast_timing,omitempty" xml:"fast_timing,omitempty"`
    // 时效类型，ESTIMATE：预估时效，PROMISE：承诺时效
    TimingType   string `json:"timing_type,omitempty" xml:"timing_type,omitempty"`
}
