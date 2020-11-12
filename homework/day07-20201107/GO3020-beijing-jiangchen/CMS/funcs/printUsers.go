package funcs

import (
	"CMS/models"
	"os"

	"github.com/olekukonko/tablewriter"
)

//PrintUsers ...
// Print specified Users slices in an elegant way
func PrintUsers(Users *[]models.User) {
	models.MUE.PersistentStorage.SyncFromDBToMemory()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"ID", "Name", "Tel", "Address", "Birthday"})
	for _, value := range *Users {
		table.Append(value.ConvertElementToSlice())
	}
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}

//PrintAllUsers ...
// Print All Users from Users slices
func PrintAllUsers() {
	models.MUE.PersistentStorage.SyncFromDBToMemory()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"ID", "Name", "Tel", "Address", "Birthday"})
	for _, value := range models.Users {
		table.Append(value.ConvertElementToSlice())
	}
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}
