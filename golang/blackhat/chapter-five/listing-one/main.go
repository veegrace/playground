package main

import (
	"github.com/miekg/dns"
)

// NOTE: Code isn't VERY USEFUL!
// Although you're sending a query to a DNS server
// and asking for the `A` record, you aren't ptocessing
// the answer.
func main() {
	// // Create a new message
	var msg dns.Msg

	// // Transforms the domain into FQDN that can be
	// // exchanged with a DNS server
	fqdn := dns.Fqdn("stacktitan.com")

	// // Modify the internal state of the message
	// // with `SetQuestion()` by using `TypeA` value
	// // to denote your intent to look up an `A` record.
	msg.SetQuestion(fqdn, dns.TypeA)

	// // This is to send the message to the provided server address,
	// which is a DNS server operated by Google in this case.
	dns.Exchange(&msg, "8.8.8.8:53")

	panic("sa")
}
