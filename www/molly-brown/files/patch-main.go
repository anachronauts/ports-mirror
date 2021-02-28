--- main.go.orig	2021-02-28 11:48:04.110000000 -0500
+++ main.go	2021-02-28 13:20:55.700000000 -0500
@@ -2,10 +2,13 @@
 
 import (
 	"crypto/tls"
+	"errors"
 	"flag"
 	"log"
 	"os"
+	"os/user"
 	"strconv"
+	"syscall"
 )
 
 func main() {
@@ -62,6 +65,19 @@
 		ClientAuth:   tls.RequestClientCert,
 	}
 
+	// Drop privileges
+	if dropped, err := dropPrivileges("gemini"); err != nil {
+		errorLog.Printf("Error dropping privileges: %v\n", err)
+		log.Fatal(err)
+	} else if dropped {
+		// Sanity check: make sure we can't read our key anymore
+		if key, err := os.Open(config.KeyPath); err == nil {
+			key.Close()
+			errorLog.Println("Key file is still readable!")
+			log.Fatal("Key file is still readable!")
+		}
+	}
+
 	// Create TLS listener
 	listener, err := tls.Listen("tcp", ":"+strconv.Itoa(config.Port), tlscfg)
 	if err != nil {
@@ -90,3 +106,47 @@
 	}
 
 }
+
+func dropPrivileges(newUser string) (bool, error) {
+	if os.Getuid() != 0 {
+		return false, nil
+	}
+
+	u, err := user.Lookup(newUser)
+	if err != nil {
+		return false, err
+	}
+
+	gs, err := u.GroupIds()
+	if err != nil {
+		return false, err
+	}
+	var gids []int
+	for _, g := range gs {
+		gid, err := strconv.Atoi(g)
+		if err != nil {
+			return false, err
+		}
+		gids = append(gids, gid)
+	}
+	if gid, err := strconv.Atoi(u.Gid); err != nil {
+		return false, err
+	} else if err := syscall.Setgroups(gids); err != nil {
+		return false, err
+	} else if err := syscall.Setgid(gid); err != nil {
+		return false, err
+	}
+
+	if uid, err := strconv.Atoi(u.Uid); err != nil {
+		return false, err
+	} else if err := syscall.Setuid(uid); err != nil {
+		return false, err
+	}
+
+	// Try to regain root privileges
+	if err := syscall.Setuid(0); err == nil {
+		return false, errors.New("successfully regained root privs")
+	}
+
+	return true, nil
+}
