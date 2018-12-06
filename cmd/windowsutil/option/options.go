// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package option

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/pflag"
)

//Config config
type Config struct {
	Debug       bool
	RunShell    string
	ServiceName string
}

//AddFlags config
func (c *Config) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.RunShell, "run", "", "Specify startup command")
	fs.StringVar(&c.ServiceName, "service-name", "", "Specify windows service name")
	fs.BoolVar(&c.Debug, "debug", false, "debug mode run ")
}

//Check check config
func (c *Config) Check() bool {
	if c.ServiceName == "" {
		logrus.Errorf("service name can not be empty")
		return false
	}
	if c.RunShell == "" {
		logrus.Errorf("run shell can not be empty")
		return false
	}
	return true
}