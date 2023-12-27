package scanner

import (
	"fmt"
	"os/exec"
)

// Scanner object
type Scanner struct {
	command  string
	enable   bool
	defaults []string
}

// Scan a host by using apt-scanner
func (s Scanner) Scan(params map[string]string) ([]string, error) {
	r := new(report)

	// load default vulnerabilities
	for _, tmp := range s.defaults {
		r.vulnerabilities = append(r.vulnerabilities, tmp)
	}

	// check scanner enable
	if !s.enable {
		return r.vulnerabilities, nil
	}

	// command flags
	flags := make([]string, 0)

	for key := range params {
		flags = append(flags, fmt.Sprintf("--%s", key), params[key])
	}

	// execute command
	cmd := exec.Command(s.command, flags...)
	if err := cmd.Start(); err != nil {
		return r.vulnerabilities, err
	}

	// read output
	context, err := cmd.Output()
	if err != nil {
		return r.vulnerabilities, err
	}

	// convert type to our report
	if er := convert(context, r); er != nil {
		return r.vulnerabilities, er
	}

	return r.vulnerabilities, nil
}
