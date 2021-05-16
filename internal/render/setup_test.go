package render

import (
	"testing"
	"encoding/gob"
	"net/http"
	"os"
	"time"
	"github.com/NamanBalaji/bookings/internal/models"
	"github.com/NamanBalaji/bookings/internal/config"
	"github.com/alexedwards/scs/v2"
)


var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M){
	//what am i going to put in the session
	gob.Register(models.Reservation{})
	// change this to true when in production
	testApp.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}