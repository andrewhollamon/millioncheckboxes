package main

import (
	"fmt"
	"github.com/andrewhollamon/millioncheckboxes/api/memoryStore"
	"github.com/andrewhollamon/millioncheckboxes/api/uuidService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func checkboxCheck(c *gin.Context) {
	doCheckboxAction(c, true)
}

func checkboxUncheck(c *gin.Context) {
	doCheckboxAction(c, false)
}

func doCheckboxAction(c *gin.Context, checked bool) {
	checkboxNbrString := strings.TrimSpace(c.Param("checkboxNbr"))
	if len(checkboxNbrString) <= 0 {
		c.String(http.StatusBadRequest, "Checkbox Number must be specified")
		return
	}
	checkboxNbr, converr := strconv.Atoi(checkboxNbrString)
	if converr != nil {
		c.String(http.StatusBadRequest, "Value for checkbox number could not be parsed as integer")
		return
	}

	requestId, uuiderr := uuidService.NewRequestUuid()
	if uuiderr != nil {
		c.String(http.StatusInternalServerError, "Could not generate Request UUID")
		return
	}
	fmt.Printf("Request [%s] to set checkbox number [%d] to %s", requestId, checkboxNbr, strconv.FormatBool(checked))

	//todo test whether queue is live, fail if not
	
	// update memory store
	checkerr := memoryStore.DoCheck(checkboxNbr, checked)
	if checkerr != nil {
		c.String(http.StatusInternalServerError, "Could not update the checkbox")
		return
	}
	fmt.Printf("Successfully updated Checkbox Number [%i] to %s", checkboxNbr, strconv.FormatBool(checked))

	// submit update to queue

	// return request id
	c.String(http.StatusOK, requestId.String())
	return
}
