# Bookings

I built this project to learn the ins and out of backend dev using golang. 

### Project functionalities
- This is room booking system`
- Room owners can log on and list their rooms
- Users can book rooms based on availability
-  After successful booking of the room a confirmation mail is sent to the user
- There's admin panel too where admin can perform housekeeping


###  Tech Used
- Golang for backend development
- HTML, CSS and JavaScript for frontend
- PostgreSQL for database
- Uses the [chi router](github.com/go-chi/chi)
- Uses [alex edwards scs session management](github.com/alexedwards/scs)
 - Uses [nosurf](github.com/justinas/nosurf) 

### Run this project locally 
- after cloning the repo run ``` go get ```  make sure you are in the root level 
- run ```soda create``` to create the database locally
- run ```soda migrate``` to create the tables
- change the ```dbname``` and ```dbuser``` in ```run.sh``` and run the command ```./run.sh```
