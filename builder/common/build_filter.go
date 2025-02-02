package common

import (
	"encoding/csv"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// Build a slice of EC2 (AMI/Subnet/VPC) filter options from the filters provided.
func buildEc2Filters(input map[string]string) []*ec2.Filter {
	var filters []*ec2.Filter

	for k, v := range input {
		var b []*string

		a := k
		csvReader := csv.NewReader(strings.NewReader(v))
		csvReader.TrimLeadingSpace = true

		values, err := csvReader.Read()
		if err != nil {
			log.Fatal(err)
		}

		for _, r := range values {
			var value = r
			b = append(b, &value)
		}

		filters = append(filters, &ec2.Filter{
			Name:   &a,
			Values: b,
		})
	}
	return filters
}
