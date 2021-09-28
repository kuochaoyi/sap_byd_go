package soap

type CustomerInvoiceRequestBundleMaintainRequestSync struct {
	Header
	Body struct {
		Text                                            string `xml:",chardata"`
		CustomerInvoiceRequestBundleMaintainRequestSync struct {
			Text                               string `xml:",chardata"`
			BusinessDocumentBasicMessageHeader string `xml:"BusinessDocumentBasicMessageHeader"`
			CustomerInvoiceRequest             struct {
				Text                              string `xml:",chardata"`
				ActionCode                        string `xml:"actionCode,attr"`
				BaseBusinessTransactionDocumentID string `xml:"BaseBusinessTransactionDocumentID"`
				BuyerParty                        struct {
					Text       string `xml:",chardata"`
					InternalID string `xml:"InternalID"`
				} `xml:"BuyerParty"`
				SalesUnitParty struct {
					Text       string `xml:",chardata"`
					InternalID string `xml:"InternalID"`
				} `xml:"SalesUnitParty"`
				PricingTerms struct {
					Text                 string `xml:",chardata"`
					PricingProcedureCode string `xml:"PricingProcedureCode"`
					CurrencyCode         string `xml:"CurrencyCode"`
				} `xml:"PricingTerms"`
				Item struct {
					Text                                        string `xml:",chardata"`
					BaseBusinessTransactionDocumentItemID       string `xml:"BaseBusinessTransactionDocumentItemID"`
					BaseBusinessTransactionDocumentItemTypeCode string `xml:"BaseBusinessTransactionDocumentItemTypeCode"`
					ReceivablesPropertyMovementDirectionCode    string `xml:"ReceivablesPropertyMovementDirectionCode"`
					Product                                     struct {
						Text       string `xml:",chardata"`
						InternalID string `xml:"InternalID"`
						TypeCode   string `xml:"TypeCode"`
					} `xml:"Product"`
					Quantity struct {
						Text     string `xml:",chardata"`
						UnitCode string `xml:"unitCode,attr"`
					} `xml:"Quantity"`
					QuantityTypeCode string `xml:"QuantityTypeCode"`
					PriceAndTax      struct {
						Text           string `xml:",chardata"`
						PriceComponent struct {
							Text     string `xml:",chardata"`
							TypeCode string `xml:"TypeCode"`
							Rate     struct {
								Text            string `xml:",chardata"`
								DecimalValue    string `xml:"DecimalValue"`
								MeasureUnitCode string `xml:"MeasureUnitCode"`
							} `xml:"Rate"`
						} `xml:"PriceComponent"`
					} `xml:"PriceAndTax"`
				} `xml:"Item"`
			} `xml:"CustomerInvoiceRequest"`
		} `xml:"CustomerInvoiceRequestBundleMaintainRequest_sync"`
	} `xml:"Body"`
}
