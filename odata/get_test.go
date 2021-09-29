package odata

import (
	"testing"
)

func TestMakeUrl(t *testing.T) {
	type args struct {
		u *Get
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "cinv",
			args: args{
				u: &Get{
					ServerCode:                 "356235",
					CustomApiName:              "cinv",
					BusinessObjectCollection:   "CustomerInvoiceCustomerInvoiceCollection",
					BusinessObjectWantToExpand: "CustomerInvoiceItem/CustomerInvoicePriceAndTaxCalculationItem",
					FilterDate:                 "2021-09-01",
				},
			},
			want: `https://my356235.sapbydesign.com/sap/byd/odata/cust/v1/cinv/CustomerInvoiceCustomerInvoiceCollection?$expand=CustomerInvoiceItem/CustomerInvoicePriceAndTaxCalculationItem&$filter=Date%20ge%20datetime%272021-09-01T00:00:00%27&$format=json`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeUrl(tt.args.u); got != tt.want {
				t.Errorf("MakeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
