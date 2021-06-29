package util

// QimenUser 
type QimenUser struct {
    // memo
    Memo   string `json:"memo,omitempty" xml:"memo,omitempty"`
    // gmtCreate
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // sellerNick
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
}
