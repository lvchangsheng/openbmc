/**
 * Copyright 2020-present Facebook. All Rights Reserved.
 *
 * This program file is free software; you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the
 * Free Software Foundation; version 2 of the License.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program in a file named COPYING; if not, write to the
 * Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor,
 * Boston, MA 02110-1301 USA
 */

package flash

import (
	"fmt"
	"log"

	"github.com/facebook/openbmc/tools/flashy/lib/step"
	"github.com/facebook/openbmc/tools/flashy/lib/utils"
	"github.com/pkg/errors"
)

func FlashFwUtil(stepParams step.StepParams) step.StepExitError {
	log.Printf("Flashing using fw-util method")
	log.Printf("Attempting to flash with image file '%v'", stepParams.ImageFilePath)

	err := runFwUtilCmd(stepParams.ImageFilePath)
	if err != nil {
		return step.ExitSafeToReboot{err}
	}
	return nil
}

var runFwUtilCmd = func(imageFilePath string) error {
	log.Printf("Starting to run fw-util")

	flashCmd := []string{
		"fw-util", "bmc", "--update", "bmc", imageFilePath,
	}

	exitCode, err, stdout, stderr := utils.RunCommand(flashCmd, 1800)
	if err != nil {
		errMsg := fmt.Sprintf(
			"Flashing failed with exit code %v, error: %v, stdout: '%v', stderr: '%v'",
			exitCode, err, stdout, stderr,
		)
		return errors.Errorf(errMsg)
	}
	log.Printf("fw-util succeeded")
	return nil
}