// Copyright (C) 2018 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package netpb

import (
	"github.com/gogo/protobuf/proto"
	"github.com/nebulasio/go-nebulas/util/logging"
	"github.com/sirupsen/logrus"
)

func HelloMessageFromProto(data []byte) (*Hello, error) {
	pb := new(Hello)

	if err := proto.Unmarshal(data, pb); err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"err": err,
		}).Debug("Invalid Hello proto message.")
		return nil, err
	}

	return pb, nil
}

func OKMessageFromProto(data []byte) (*OK, error) {
	pb := new(OK)

	if err := proto.Unmarshal(data, pb); err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"err": err,
		}).Debug("Invalid OK proto message.")
		return nil, err
	}

	return pb, nil
}
