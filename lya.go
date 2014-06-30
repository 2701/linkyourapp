package lya

import (
	"errors"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

type AppLink struct {
	AppLinkId   string
	DefaultLink string
	IOSLink     string
	AndroidLink string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", root)
	http.HandleFunc("/add", add)
}

// appLinkKey returns a new Key to be used for a new AppLink
func appLinkKey(c appengine.Context) *datastore.Key {
	keyId := NewId()
	return datastore.NewKey(c, "AppLink", keyId, 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	q := datastore.NewQuery("AppLink").Limit(10)
	appLinks := make([]AppLink, 0, 10)
	appLinkKeys, err := q.GetAll(c, &appLinks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Template data
	data := struct {
		AppLinks []AppLink
		Keys     []*datastore.Key
	}{
		appLinks,
		appLinkKeys,
	}
	err = appLinkTemplates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var appLinkTemplates = template.Must(template.ParseFiles("templates/index.html"))

// var appLinkTemplate = template.Must(template.New("appLink").Parse(`
// <html>
//   <head>
//     <title>AppLinks!</title>
//   </head>
//   <body>
//  {{range $i, $appLink := .AppLinks}}
//    <ul>{{index $.Keys $i}}
//      <li>Default Link: {{$appLink.DefaultLink}}</li>
//    	<li>iOS Link: {{$appLink.IOSLink}}</li>
//    	<li>Android Link: {{$appLink.AndroidLink}}</li>
//    </ul>
// {{else}}
// 	<p>No links in the datastore yet</p>
//  {{end}}
//     <form action="/add" method="post">
//       <div><label>Default Link:</label> <input type="text" name="DefaultLink" size="60"></input></div>
//       <div><label>iOS Link:</label> <input type="text" name="IOSLink" size="60"></input></div>
//       <div><label>Android Link:</label> <input type="text" name="AndroidLink" size="60"></input></div>
//       <div><input type="submit" value="AddLink"></div>
//     </form>
//   </body>
// </html>
// `))

func add(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// TODO: We could do this validation with https://github.com/gorilla/schema
	switch {
	case r.FormValue("DefaultLink") == "":
		err := errors.New("Please provide links to all platforms")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case r.FormValue("IOSLink") == "":
		err := errors.New("Please provide links to all platforms")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case r.FormValue("AndroidLink") == "":
		err := errors.New("Please provide links to all platforms")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a := AppLink{
		DefaultLink: r.FormValue("DefaultLink"),
		IOSLink:     r.FormValue("IOSLink"),
		AndroidLink: r.FormValue("AndroidLink"),
	}

	// We set the same parent key on every AppLink entity to ensure each AppLink
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := appLinkKey(c)
	_, err := datastore.Put(c, key, &a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func validateUrl(u string) error {

	return errors.New("We don't know what's that")
}
