package main

import "sync"

func addSystemUser(userName string, password string, wg *sync.WaitGroup) bool {
	runcmd("useradd -m -p " + password + " " + userName)
	wg.Done()
	return false
}
func addMysqlUser(userName string, password string, wg *sync.WaitGroup) bool {
	runcmd(`mysql --login-path=adci -e "CREATE USER ` + userName + `@localhost IDENTIFIED BY '` + password + `';"`)
	runcmd(`mysql --login-path=adci -e "GRANT ALL PRIVILEGES ON *.* TO ` + userName + `@localhost IDENTIFIED BY '` + password + `';"`)
	runcmd(`mysql --login-path=adci -e "FLUSH PRIVILEGES;"`)
	wg.Done()
	return false
}
func deleteSystemUser(userName string, wg *sync.WaitGroup) bool {
	runcmd("useradd  " + userName)
	wg.Done()
	return false
}
func deleteMysqlUser(userName string, wg *sync.WaitGroup) bool {
	// runcmd(`mysql -uroot -prootpw -e "CREATE USER ` + userName + `@localhost IDENTIFIED BY '` + password + `';"`)
	// runcmd(`mysql -uroot -prootpw -e "GRANT ALL PRIVILEGES ON *.* TO ` + userName + `@localhost IDENTIFIED BY '` + password + `';"`)
	// runcmd(`mysql -uroot -prootpw -e "FLUSH PRIVILEGES;"`)
	wg.Done()
	return false
}
func addUser(userName string) bool {
	var password = randomString(8)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go addSystemUser(userName, password, wg)
	go addMysqlUser(userName, password, wg)
	wg.Wait()

	return false
}
func deleteUser(userName string) bool {
	// sword = randomString(8)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go deleteSystemUser(userName, wg)
	go deleteMysqlUser(userName, wg)
	wg.Wait()
	//@TODO print user banner
	return false
}
