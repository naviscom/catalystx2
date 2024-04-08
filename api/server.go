package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/naviscom/catalystx2/db/sqlc"
	"github.com/naviscom/catalystx2/token"
	"github.com/naviscom/catalystx2/util"
)

// // Server serves HTTP requests.
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP Gin server and set up routing & CORS.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//router.POST("/users", server.createUser)
	//router.POST("/users/login", server.loginUser)
	//router.POST("/tokens/renew_access", server.renewAccessToken)
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/serviceareatypes", server.createServiceareatype)
	authRoutes.GET("/serviceareatypes0/:id", server.getServiceareatype0)
	authRoutes.GET("/serviceareatypes1/:serviceareatype_name", server.getServiceareatype1)
	authRoutes.GET("/serviceareatypes", server.listServiceareatypes)
	authRoutes.POST("/updateserviceareatypes", server.updateServiceareatype)
	authRoutes.GET("/deleteserviceareatypes/:id", server.deleteServiceareatype)
	authRoutes.POST("/areas", server.createArea)
	authRoutes.GET("/areas0/:id", server.getArea0)
	authRoutes.GET("/areas1/:area_name", server.getArea1)
	authRoutes.GET("/areas", server.listAreas)
	authRoutes.POST("/updateareas", server.updateArea)
	authRoutes.GET("/deleteareas/:id", server.deleteArea)
	authRoutes.POST("/clutters", server.createClutter)
	authRoutes.GET("/clutters0/:id", server.getClutter0)
	authRoutes.GET("/clutters1/:clutter_name", server.getClutter1)
	authRoutes.GET("/clutters", server.listClutters)
	authRoutes.POST("/updateclutters", server.updateClutter)
	authRoutes.GET("/deleteclutters/:id", server.deleteClutter)
	authRoutes.POST("/sitetypes", server.createSitetype)
	authRoutes.GET("/sitetypes0/:id", server.getSitetype0)
	authRoutes.GET("/sitetypes1/:type_name", server.getSitetype1)
	authRoutes.GET("/sitetypes", server.listSitetypes)
	authRoutes.POST("/updatesitetypes", server.updateSitetype)
	authRoutes.GET("/deletesitetypes/:id", server.deleteSitetype)
	authRoutes.POST("/vendors", server.createVendor)
	authRoutes.GET("/vendors0/:id", server.getVendor0)
	authRoutes.GET("/vendors1/:vendor_name", server.getVendor1)
	authRoutes.GET("/vendors", server.listVendors)
	authRoutes.POST("/updatevendors", server.updateVendor)
	authRoutes.GET("/deletevendors/:id", server.deleteVendor)
	authRoutes.POST("/techs", server.createTech)
	authRoutes.GET("/techs0/:id", server.getTech0)
	authRoutes.GET("/techs1/:tech_name", server.getTech1)
	authRoutes.GET("/techs", server.listTechs)
	authRoutes.POST("/updatetechs", server.updateTech)
	authRoutes.GET("/deletetechs/:id", server.deleteTech)
	authRoutes.POST("/bands", server.createBand)
	authRoutes.GET("/bands0/:id", server.getBand0)
	authRoutes.GET("/bands1/:band_name", server.getBand1)
	authRoutes.GET("/bands", server.listBands)
	authRoutes.POST("/updatebands", server.updateBand)
	authRoutes.GET("/deletebands/:id", server.deleteBand)
	authRoutes.POST("/carriers", server.createCarrier)
	authRoutes.GET("/carriers0/:id", server.getCarrier0)
	authRoutes.GET("/carriers1/:carrier_name", server.getCarrier1)
	authRoutes.GET("/carriers", server.listCarriers)
	authRoutes.POST("/updatecarriers", server.updateCarrier)
	authRoutes.GET("/deletecarriers/:id", server.deleteCarrier)
	authRoutes.POST("/continents", server.createContinent)
	authRoutes.GET("/continents0/:id", server.getContinent0)
	authRoutes.GET("/continents1/:continent_name", server.getContinent1)
	authRoutes.GET("/continents", server.listContinents)
	authRoutes.POST("/updatecontinents", server.updateContinent)
	authRoutes.GET("/deletecontinents/:id", server.deleteContinent)
	authRoutes.POST("/countries", server.createCountry)
	authRoutes.GET("/countries0/:id", server.getCountry0)
	authRoutes.GET("/countries1/:country_name", server.getCountry1)
	authRoutes.GET("/countries", server.listCountries)
	authRoutes.POST("/updatecountries", server.updateCountry)
	authRoutes.GET("/deletecountries/:id", server.deleteCountry)
	authRoutes.POST("/states", server.createState)
	authRoutes.GET("/states0/:id", server.getState0)
	authRoutes.GET("/states1/:state_name", server.getState1)
	authRoutes.GET("/states", server.listStates)
	authRoutes.POST("/updatestates", server.updateState)
	authRoutes.GET("/deletestates/:id", server.deleteState)
	authRoutes.POST("/cities", server.createCity)
	authRoutes.GET("/cities0/:id", server.getCity0)
	authRoutes.GET("/cities1/:city_name", server.getCity1)
	authRoutes.GET("/cities", server.listCities)
	authRoutes.POST("/updatecities", server.updateCity)
	authRoutes.GET("/deletecities/:id", server.deleteCity)
	authRoutes.POST("/districts", server.createDistrict)
	authRoutes.GET("/districts0/:id", server.getDistrict0)
	authRoutes.GET("/districts1/:district_name", server.getDistrict1)
	authRoutes.GET("/districts", server.listDistricts)
	authRoutes.POST("/updatedistricts", server.updateDistrict)
	authRoutes.GET("/deletedistricts/:id", server.deleteDistrict)
	authRoutes.POST("/towns", server.createTown)
	authRoutes.GET("/towns0/:id", server.getTown0)
	authRoutes.GET("/towns1/:town_name", server.getTown1)
	authRoutes.GET("/towns", server.listTowns)
	authRoutes.POST("/updatetowns", server.updateTown)
	authRoutes.GET("/deletetowns/:id", server.deleteTown)
	authRoutes.POST("/blocks", server.createBlock)
	authRoutes.GET("/blocks0/:id", server.getBlock0)
	authRoutes.GET("/blocks1/:block_name", server.getBlock1)
	authRoutes.GET("/blocks", server.listBlocks)
	authRoutes.POST("/updateblocks", server.updateBlock)
	authRoutes.GET("/deleteblocks/:id", server.deleteBlock)
	authRoutes.POST("/properties", server.createProperty)
	authRoutes.GET("/properties0/:id", server.getProperty0)
	authRoutes.GET("/properties1/:property_name", server.getProperty1)
	authRoutes.GET("/properties", server.listProperties)
	authRoutes.POST("/updateproperties", server.updateProperty)
	authRoutes.GET("/deleteproperties/:id", server.deleteProperty)
	authRoutes.POST("/sites", server.createSite)
	authRoutes.GET("/sites0/:id", server.getSite0)
	authRoutes.GET("/sites1/:site_name", server.getSite1)
	authRoutes.GET("/sites", server.listSites)
	authRoutes.POST("/updatesites", server.updateSite)
	authRoutes.GET("/deletesites/:id", server.deleteSite)
	authRoutes.POST("/cells", server.createCell)
	authRoutes.GET("/cells0/:id", server.getCell0)
	authRoutes.GET("/cells1/:cell_name", server.getCell1)
	authRoutes.GET("/cells", server.listCells)
	authRoutes.POST("/updatecells", server.updateCell)
	authRoutes.GET("/deletecells/:id", server.deleteCell)
	authRoutes.POST("/traffic", server.createTraffic)
	authRoutes.GET("/traffic0/:id", server.getTraffic0)
	authRoutes.GET("/traffic", server.listTraffic)
	authRoutes.POST("/updatetraffic", server.updateTraffic)
	authRoutes.GET("/deletetraffic/:id", server.deleteTraffic)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.GET("/users0/:username", server.getUser0)
	router.GET("/users3/:email", server.getUser3)
	router.GET("/users", server.listUsers)
	router.POST("/updateusers", server.updateUser)
	router.GET("/deleteusers/:id", server.deleteUser)
	authRoutes.POST("/sessions", server.createSession)
	authRoutes.GET("/sessions0/:id", server.getSession0)
	authRoutes.GET("/sessions", server.listSessions)
	authRoutes.POST("/updatesessions", server.updateSession)
	authRoutes.GET("/deletesessions/:id", server.deleteSession)

	// setting the Gin router to Default() allows all origins
	router.Use(cors.Default())

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
