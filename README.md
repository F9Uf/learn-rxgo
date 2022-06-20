# Learn RXGO

There is a list of `Customer`, and each `Customer` has a structure below:
```go
type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}
```
Our goal is need to perform the two following oerpations:
- filter the customers whose age is below 18
- enrich each customer with a tax number

