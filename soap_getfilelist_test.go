package expertview

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseGetFileList(t *testing.T) {
	r := `<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <ns2:getFileListResponse xmlns:ns2="http://webservice.expertview.squarell.com/">
         <return>PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+PFJFU1BPTlNFPjxERVZJQ0VUWVBFUz48REVWSUNFVFlQRSBERVNDUklQVElPTj0iTm9ybWFsIChTb2xpZC9GbGV4KSIgUFJPRFVDVE5VTUJFUj0iODAwMC0xIi8+PERFVklDRVRZUEUgREVTQ1JJUFRJT049IkJsdWV0b290aCAyLjEiIFBST0RVQ1ROVU1CRVI9Ijg1OTAtMTAiLz48REVWSUNFVFlQRSBERVNDUklQVElPTj0iREFQIiBQUk9EVUNUTlVNQkVSPSI2NjM0LTY3MCIvPjwvREVWSUNFVFlQRVM+PEZJTEVTPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5JTlRBMS1GTFgxMi1NQkErVkRPLVJTLTEyMDIyNENMLkRDRjwvTkFNRT48RklMRT5EOTk4NDUyNzU4MjAxMjAyMjcxNTE4NDc0Ny5kY2Y8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLVJFTEFZLTAwMDAwMC0xMjEwMTIuRENGPC9OQU1FPjxGSUxFPkQ5MDkzNDc1MzQyMDEyMTAxMjE1MDcwMjU1LmRjZjwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLUZMWDEyLVRESy0xMjExMTNDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMzczMzgwNTk4MjQ2MTE2Njg3LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+VEVTVC04MDAwLVRSQUlMRVJJRC0xMzAyMThDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNjIxMjA0NjAwNzcyMTk3MTIwMi5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPkFUUi04MDAwLUJNVy0wMDAwMDAtMTMwMzEzQ0wuRENGPC9OQU1FPjxGSUxFPkRTUTU1MjEzMjg2NzY0NDg2Nzk0NTIuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1PQkQtMDAwMDAwLTEzMDgwMkNMLkRDRjwvTkFNRT48RklMRT5EU1EyNzQ3NDY4MTQ1MjE1NDI4NTg0LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtSVNVLTAwMDAwMC0xMzEwMDFDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNjI1Nzk3OTU2OTY1Mjk0MDQ4MS5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPklOVC04MDAwLVBTQi1CRVRBMDAtMTQwNzE3Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTc1MDU3MTE0NzQ1NTMwNTM5NjYuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5NVVctODAwMC1QU0ItQkVUQTAwLTE0MDcxN0NMLkRDRjwvTkFNRT48RklMRT5EU1EyNjMyMTk5NDI1MDIxMTA4NTY1LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtT0JELUJFVEEwMC0xNDA5MDFDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNjYwMTc1NzYxMTM2NjE1ODkxNy5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLU9CRC1CRVRBMDAtNDMwTk0tMTQwOTAxQ0wuRENGPC9OQU1FPjxGSUxFPkRTUTQ1ODcxMzA5NTk2OTUyMjM2MTguZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5BVFItODAwMC1ST04tQkVUQTAwLTE1MDExMkNMLkRDRjwvTkFNRT48RklMRT5EU1ExNzEwMzY5NjY5NTYyODEzMjIxLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+QVRSLTgwMDAtVE9ZLTAwMDAwMC0xNTAyMjZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRODcxNzA2MjI3NTgwMzkzMDUyOS5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPk1VVy04MDAwLVBTQi1CRVRBMDAtMTUwMzI0Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTU1MjI4Njg4Njc4MTMwNDgyNC5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLU1CUy1UUFQwMDAtQlQrRFVQLTE1MDcwOUNMLkRDRjwvTkFNRT48RklMRT5EU1E2ODc5MzYzMTY1Mzc4NTUzMTg4LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtQ09OLTI1NkstRlc0OS0wMDAwMDAtMTUwODI4Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTczOTMxMjMyNDE0Mjk2MDcxMDIuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5BVFItODAwMC1QU0EtMjU2Sy1GVzQ5LUJFVEEwMC0xNTExMTNDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMjI2NTU4NTM2NjAwNTQ1MDg1OS5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlJFU0VULUZXNDgrRFVQLTE2MDEyNkNMLkRDRjwvTkFNRT48RklMRT5EU1EyMDc1ODgyODAxODY3MzQ5Mjg2LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtU1VCLTI1NkstRlc0OS0wMDAwMDAtMTYwMTI4Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTYxMjAzODQ0Nzk1MTQyNjcyMzIuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5JTlQtODAwMC1UUktTLTI1NkstRlc0OS0wMDAwMDAtMTUxMjAzQ0wuRENGPC9OQU1FPjxGSUxFPkRTUTM2MDQ1Mzk4NzQwNTU2NDY4MjYuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5JTlQtODAwMC1SRU4tMjU2Sy1GVzQ5LUJFVEEwMC0xNjA1MTdDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNDU2NjEyNzczMDU1NTYxNzI3My5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLUJNVy0yNTZLLUZXNDktMDAwMDAwLTE2MDYxNkNMLkRDRjwvTkFNRT48RklMRT5EU1ExNDk1ODM0MTA3Mzk3MDY2NzgxLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtRlRDLTI1NkstRlc0OS0wMDAwMDAtMTYwNjE2Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTMwODE5ODk1MzU3OTIxNTc5NjQuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1GVE0tMjU2Sy1GVzQ5LTAwMDAwMC0xNjA2MTZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNzg3ODcxMzk0Njc5ODY0ODM4My5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLUZUUy0yNTZLLUZXNDktMDAwMDAwLTE2MDYxNkNMLkRDRjwvTkFNRT48RklMRT5EU1E1MzE3MDQxMTI5NjUxMjcxNTM4LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtRlVBLTI1NkstRlc0OS0wMDAwMDAtMTYwNjE2Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTQ5ODAyMzQyOTM2NDMxNzcwMzUuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1ISU5PLTI1NkstRlc0OS0wMDAwMDAtMTYwNjE2Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTQyMzA5MTA2NTkxNTQzMDMxNTYuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1IT04tMjU2Sy1GVzQ5LTAwMDAwMC0xNjA2MTZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMTE4NTk5NDY1OTQxMTg4MDY5Mi5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLUhZVS0yNTZLLUZXNDktMDAwMDAwLTE2MDYxNkNMLkRDRjwvTkFNRT48RklMRT5EU1E3Njg4ODQxNTc3MzE4NjA1NDg2LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtSkFHLTI1NkstRlc0OS0wMDAwMDAtMTYwNjE2Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTI3NzczNDAyNjE3OTkzMTc4MDcuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1NQVotMjU2Sy1GVzQ5LTAwMDAwMC0xNjA2MTZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRODA1MjkxNjE1MjgwNzYzNTg0Ni5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLU1JUC0yNTZLLUZXNDktMDAwMDAwLTE2MDYxNkNMLkRDRjwvTkFNRT48RklMRT5EU1E3MjgzMjMyMDM4MzYzOTAwMDk3LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtT1BMLTI1NkstRlc0OS0wMDAwMDAtMTYwNjE2Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTExMDM4NDU3MjQwNjk1MDg1MDkuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1QU0ItMjU2Sy1GVzQ5LTAwMDAwMC0xNjA2MTZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMjE1MTczODM0NzM4MDk0MTguZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1SRU4tMjU2Sy1GVzQ5LTAwMDAwMC0xNjA2MTZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMjc5Njg0MDEyNTc3NzY4MjIwNi5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLVJPTi0yNTZLLUZXNDktMDAwMDAwLTE2MDYxNkNMLkRDRjwvTkFNRT48RklMRT5EU1E1MDYwODgzMzMzMjgzMTY1MzQ4LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtVE9ZLTI1NkstRlc0OS0wMDAwMDAtMTYwNjE3Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTI0OTAzMzA1MTg5MDg0MzgxNTYuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1WQUctMjU2Sy1GVzQ5LTAwMDAwMC0xNjA2MTZDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNTE1NTgzNzM2ODQ4ODM5NTE4OC5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLVZWVi0yNTZLLUZXNDktMDAwMDAwLTE2MDYxNkNMLkRDRjwvTkFNRT48RklMRT5EU1EzMTY1MTgzODI5MTI2OTU3MTMyLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtVFJLUy0yNTZLLUZXNDktMDAwMDAwLTE2MDYyMUNMLkRDRjwvTkFNRT48RklMRT5EU1E2MzM4OTMzNDM4MDkyOTE4NjkxLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+U1FVLTgwMDAtVFJLUy0yNTZLLUZXNDktRk1TMDAwLTE2MDYyMUNMLkRDRjwvTkFNRT48RklMRT5EU1E0ODM5MjcxNDU2NjI1MzUwNDguZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1UUktTLTI1NkstRlc0OS1UUFQwMDAtMTYwNjIxQ0wuRENGPC9OQU1FPjxGSUxFPkRTUTc3MzE0NjkwOTE3OTQ0NTE3MDcuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5NVVctODAwMC1UUktTLTI1NkstRlc0OS0wMDAwMDAtMTYwNzA3Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTg4NTMzNjUyMzMyMDQ0MzQyNTIuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1QU0EtMjU2Sy1GVzQ5LTAwMDAwMC0xNjA3MTFDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNTEyNTY1MDY3MDI2MTM3NjgyOS5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPkFUUi04MDAwLVRSS1MtMjU2Sy1GVzQ5LTAwMDAwMC0xNjA3MDVDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNDE4OTM0NDIzMDIwMzczMjg0NC5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPkFUUi04MDAwLVZBRy0yNTZLLUZXNDktMDAwMDAwLTE2MDcyOENMLkRDRjwvTkFNRT48RklMRT5EU1E3NzM2NzE5ODEyNzIwOTY4Njg4LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+QVRSLTgwMDAtUkVOLTI1NkstRlc0OS0wMDAwMDAtMTYwNzI4Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTEzODAxMzg3MTM3NzQ2MTUzNDQuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5BVFItODAwMC1QU0ItMjU2Sy1GVzQ5LTAwMDAwMC0xNjA3MjhDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRNDc3ODIwMTY3NzQ1OTk4MDA2Ny5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPkFUUi04MDAwLU9QTC0yNTZLLUZXNDktMDAwMDAwLTE2MDcyOENMLkRDRjwvTkFNRT48RklMRT5EU1E3MTI4NTcyNDY5NzU0NDA1NDY1LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkRDRiI+PE5BTUU+QVRSLTgwMDAtTUJTLTI1NkstRlc0OS0wMDAwMDAtMTYwNzI4Q0wuRENGPC9OQU1FPjxGSUxFPkRTUTU0NjEwNzY1NzQwMDk1MDI0OTYuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5BVFItODAwMC1GVFMtMjU2Sy1GVzQ5LTAwMDAwMC0xNjA3MjhDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMTg5NTAyNDc2NjEwNTQ4MTY3Ny5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJEQ0YiPjxOQU1FPlNRVS04MDAwLUNPTi0yNTZLLUZXNDktMDAwMDAwLTE2MDgwNUNMLkRDRjwvTkFNRT48RklMRT5EU1E5NTQ3MTMyNzAwODYwOTA0MzMuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRENGIj48TkFNRT5TUVUtODAwMC1NQlMtMjU2Sy1GVzQ5LTAwMDAwMC0xNjA4MjJDTC5EQ0Y8L05BTUU+PEZJTEU+RFNRMzQyMDIyNjgyOTY3MjU4MjgwLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkZpcm13YXJlIiBQUk9EVUNUTlVNQkVSPSI4MDAwLTEiPjxOQU1FPjgwMDAtMDFWMTAxUjA0MS5CSU48L05BTUU+PEZJTEU+RlNRNDg3NzAxNTgxMjY2MzEzODgzLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkZpcm13YXJlIiBQUk9EVUNUTlVNQkVSPSI4MDAwLTEiPjxOQU1FPjgwMDAtMDFWMTEwUjA0Mi5CSU48L05BTUU+PEZJTEU+RlNRMjg0NjA4Njc1NDA3ODM4NzE3NC5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PFJFQ09SRCBLSU5EPSJGaXJtd2FyZSIgUFJPRFVDVE5VTUJFUj0iODAwMC0xIj48TkFNRT44MDAwLTAxVjExMVIwNDMuQklOPC9OQU1FPjxGSUxFPkZTUTc0ODQ4OTgyMDEwMzIwNTU1MzAuZGF0PC9GSUxFPjwvUkVDT1JEPjxSRUNPUkQgS0lORD0iRmlybXdhcmUiIFBST0RVQ1ROVU1CRVI9IjgwMDAtMSI+PE5BTUU+ODAwMC0wMVYxMTRSMDQ4LkJJTjwvTkFNRT48RklMRT5GU1ExNDk5OTA1NzM0MzAwOTM1ODQ4LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkZpcm13YXJlIiBQUk9EVUNUTlVNQkVSPSI4MDAwLTEiPjxOQU1FPjgwMDAtMDFWMTE1UjA0OS5CSU48L05BTUU+PEZJTEU+RlNROTA1NTE0NzIxMDAxOTQ1MDA1LmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkZpcm13YXJlIiBQUk9EVUNUTlVNQkVSPSI4NTkwLTEwIj48TkFNRT44NTkwSDFWMTM1LkJJTjwvTkFNRT48RklMRT5TU1EyNjEzMjU2NjgyMDc3ODU1NTUzLmRhdDwvRklMRT48L1JFQ09SRD48UkVDT1JEIEtJTkQ9IkZpcm13YXJlIiBQUk9EVUNUTlVNQkVSPSI2NjM0LTY3MCI+PE5BTUU+NjYzNEgyVjIwNC5CSU48L05BTUU+PEZJTEU+U1NRMjUzODk0NTY4Mzg4NTQ4NDUyNC5kYXQ8L0ZJTEU+PC9SRUNPUkQ+PC9GSUxFUz48L1JFU1BPTlNFPg==</return>
      </ns2:getFileListResponse>
   </S:Body>
</S:Envelope>`
	fl, err := parseGetFileList([]byte(r))
	require.Nil(t, err)
	assert.Len(t, fl.DeviceTypes, 3)
	assert.Equal(t, "8000-1", fl.DeviceTypes[0].ProductNumber)
	assert.Equal(t, "Normal (Solid/Flex)", fl.DeviceTypes[0].Description)
	assert.Len(t, fl.Records, 60)
	assert.Equal(t, DCFKind, fl.Records[0].Kind)
	assert.Equal(t, "INTA1-FLX12-MBA+VDO-RS-120224CL.DCF", fl.Records[0].Name)
	assert.Equal(t, "D9984527582012022715184747.dcf", fl.Records[0].File)
}

func TestParseGetFileListAuthentication(t *testing.T) {
	r := `<S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
   <S:Body>
      <S:Fault xmlns:ns4="http://www.w3.org/2003/05/soap-envelope">
         <faultcode>S:Server</faultcode>
         <faultstring>[2016/08/24 18:20:08.072] Login failed</faultstring>
         <detail>
            <ns2:AuthenticationException xmlns:ns2="http://webservice.expertview.squarell.com/">
               <message>[2016/08/24 18:20:08.072] Login failed</message>
            </ns2:AuthenticationException>
         </detail>
      </S:Fault>
   </S:Body>
</S:Envelope>`
	_, err := parseGetFileList([]byte(r))
	assert.Equal(t, ErrAuthentication, err)
}