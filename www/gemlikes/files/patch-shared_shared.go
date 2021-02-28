--- shared/shared.go.orig	2021-02-28 15:12:15.000000000 -0500
+++ shared/shared.go	2021-02-28 18:15:27.494922000 -0500
@@ -205,6 +205,10 @@
 }
 
 func getConfigDir() string {
+	if configDir := os.Getenv("GEMLIKES_CONFIG_DIR"); configDir != "" {
+		abs, _ := filepath.Abs(configDir)
+		return abs
+	}
 	e, _ := os.Executable()
 	return filepath.Dir(e)
 }
