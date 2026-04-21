# Backend Web Service for Notes Calendar with Go (+Gin)

## About the project

In development...

Notes-calendar-ws represents a backend web service implementation in Golang (+Gin web framework) for the <a href="https://github.com/atanas3angelov/notes-calendar"> Notes-calendar project</a>. It provides the capability to save the notes from the calendar to a database (MongoDB) so that they can be loaded, updated or analyzed. It also serves as a preview of how one can integrate the Notes-calendar into their webpage.

## For developers

### Submodules

nodes-calendar is a submodule - don't forget to pass <mark style="background-color: #ccc">--recurse-submodules</mark> to <mark style="background-color:#ccc">git clone</mark>.  
  
If you did forget to do that the notes-calendar directory will be empty and you'll need to:  
- run 2 commands from the main project: <mark style="background-color: #ccc">git submodule init</mark> to initialize your local configuration file, and <mark style="background-color: #ccc">git submodule update</mark> to fetch all the data.  
- run <mark style="background-color: #ccc">git submodule update --init</mark> or <mark style="background-color: #ccc">git submodule update --init --recursive</mark>  
  
Use <mark style="background-color: #ccc">git submodule update --remote</mark> to get the notes-calendar submodule up to date (thus locking it for the rest of the team). Then others need to run <mark style="background-color: #ccc">git submodule update --init --recursive</mark> after a <mark style="background-color: #ccc">git pull</mark> or simply <mark style="background-color: #ccc">git pull --recurse-submodules</mark>.

Read more at [Git - Submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules).  

