package ticket

// TicketSimpleSkuParam 
type TicketSimpleSkuParam struct {

    // 门票 票种类型
    
    TicketType   string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
    

    // 该票种下使用的价格规则
    
    PriceRules   []PriceRule `json:"price_rules,omitempty" xml:"price_rules,omitempty"`
    

    // 门票场次（场次门票专用）
    
    TicketSeason   string `json:"ticket_season,omitempty" xml:"ticket_season,omitempty"`
    

    // 门票区域（场次门票专用）
    
    TicketArea   string `json:"ticket_area,omitempty" xml:"ticket_area,omitempty"`
    

}
