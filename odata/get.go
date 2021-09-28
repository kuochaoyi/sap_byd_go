package odata

type Get struct {
	ServerCode string `json:"server_code"`
	// ODataName
	CustomApiName              string `json:"custom_api_name"`
	BusinessObjectCollection   string `json:"business_object_collection"`
	BusinessObjectWantToExpand string `json:"business_object_want_to_expand"`
	// 2021-09-01
	FilterDate string `json:"filter_date"`
	//Filter                string
	//Select                []string
	//OrderBy               string
	//Format                string // json
}

func MakeUrl(u *Get) string {
	s := "https://my" + u.ServerCode + ".sapbydesign.com/sap/byd/odata/cust/v1/"

	if u.CustomApiName != "" {
		s += u.CustomApiName + "/"
	}

	s += u.BusinessObjectCollection

	if u.BusinessObjectWantToExpand != "" {
		s += "?$expand=" + u.BusinessObjectWantToExpand
	}

	// rules (Optional)
	if u.FilterDate != "" {
		s += "&$filter=Date%20ge%20datetime%27" + u.FilterDate + "T00:00:00%27"
	}
	// ending
	s += "&$format=json"
	//strings.TrimRight(s, "%")
	return s
}

//https://{ServerCode}.sapbydesign.com/sap/byd/odata/cust/v1/
//{ODataName}/{BusinessObjectCollection}?
//${expand=BusinessObjectWantToExpand}&${select=aaa,bbb,ccc}&${orderby=Date%20desc}&${filter=ID%20eq%20%27999%27}&${format=json}

//GET {{SAPByD}}/cust/v1/cinv/CustomerInvoiceCustomerInvoiceCollection?$expand=CustomerInvoiceItem/CustomerInvoicePriceAndTaxCalculationItem{{format}}{{date1}}
