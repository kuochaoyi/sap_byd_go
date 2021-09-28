package soap

import "encoding/xml"

type Header struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Glob    string   `xml:"glob,attr"`
	A3k     string   `xml:"a3k,attr"`
	Glob1   string   `xml:"glob1,attr"`
	Header  string   `xml:"Header"`
}
