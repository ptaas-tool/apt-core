package scanner

import "os/exec"

// Scan a host by using apt-scanner
func Scan(command string, enable bool, defaults ...string) ([]string, error) {
	r := new(report)

	// load default vulnerabilities
	for _, tmp := range defaults {
		r.vulnerabilities = append(r.vulnerabilities, tmp)
	}

	// check scanner enable
	if !enable {
		return r.vulnerabilities, nil
	}

	// execute command
	cmd := exec.Command(command)
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
