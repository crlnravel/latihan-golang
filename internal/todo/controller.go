package todo

//
//// Controller defines the interface for exam-related HTTP handlers.
//type Controller interface {
//	GetExams(c *fiber.Ctx) error
//	GetExamByID(c *fiber.Ctx) error
//	AddExam(c *fiber.Ctx) error
//	UpdateExam(c *fiber.Ctx) error
//	DeleteExam(c *fiber.Ctx) error
//}
//
//// controller implements the Controller interface.
//type controller struct {
//	service Service
//}
//
//// NewController initializes and returns a new controller instance.
//func NewController(s Service) Controller {
//	return &controller{
//		service: s,
//	}
//}
//
//// GetExams godoc
//// @Summary      Get all exams
//// @Description  Retrieves all exams from the database.
//// @Tags         Exam
//// @Accept       json
//// @Produce      json
//// @Success      200  {object}  []entities.Exam
//// @Failure		 500  {object} ErrorResponse
//// @Router       /api/v1/tryout/exam/ [get]
//// GetExams fetches and returns all exams.
//func (ctr *controller) GetExams(c *fiber.Ctx) error {
//	fetched, err := ctr.service.FetchExams()
//	if err != nil {
//		c.Status(http.StatusInternalServerError)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	return c.JSON(fetched)
//}
//
//// GetExamByID godoc
//// @Summary      Get an exam by ID
//// @Description  Retrieves an exam by its ID.
//// @Tags         Exam
//// @Accept       json
//// @Produce      json
//// @Param        id   path      string  true  "Exam ID"
//// @Success      200  {object}  entities.Exam
//// @Failure		 500  {object} ErrorResponse
//// @Router       /api/v1/tryout/exam/{id} [get]
//// GetExamByID fetches and returns a specific exam by its ID.
//func (ctr *controller) GetExamByID(c *fiber.Ctx) error {
//	id := c.Params("id")
//	fetched, err := ctr.service.FetchExamByID(id)
//	if err != nil {
//		c.Status(http.StatusInternalServerError)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	return c.JSON(fetched)
//}
//
//// AddExam godoc
//// @Summary      Add a new exam
//// @Description  Creates a new exam in the database.
//// @Tags         Exam
//// @Accept       json
//// @Produce      json
//// @Param        exam  body      entities.Exam  true  "Exam data"
//// @Success      200   {object}  entities.Exam
//// @Failure		 400  {object} ErrorResponse
//// @Failure		 500  {object} ErrorResponse
//// @Router       /api/v1/tryout/exam/ [post]
//// AddExam creates a new exam.
//func (ctr *controller) AddExam(c *fiber.Ctx) error {
//	var body entities.Exam
//	err := c.BodyParser(&body)
//	if err != nil {
//		c.Status(http.StatusBadRequest)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	if err := ctr.service.InsertExam(&body); err != nil {
//		c.Status(http.StatusInternalServerError)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	return c.JSON(&body)
//}
//
//// UpdateExam godoc
//// @Summary      Update an existing exam
//// @Description  Updates an existing exam in the database.
//// @Tags         Exam
//// @Accept       json
//// @Produce      json
//// @Param        exam  body      entities.Exam  true  "Updated exam data"
//// @Success      200   {object}  entities.Exam
//// @Failure		 400  {object} ErrorResponse
//// @Failure		 500  {object} ErrorResponse
//// @Router       /api/v1/tryout/exam/ [put]
//// UpdateExam updates an existing exam.
//func (ctr *controller) UpdateExam(c *fiber.Ctx) error {
//	var body entities.Exam
//	err := c.BodyParser(&body)
//	if err != nil {
//		c.Status(http.StatusBadRequest)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	if err := ctr.service.UpdateExam(&body); err != nil {
//		c.Status(http.StatusInternalServerError)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	return c.JSON(&body)
//}
//
//// DeleteExam godoc
//// @Summary      Delete an exam
//// @Description  Removes an exam by its ID.
//// @Tags         Exam
//// @Accept       json
//// @Produce      json
//// @Param        id    path      string  true  "Exam ID"
//// @Success      200   {object}  string  "Exam ID"
//// @Failure		 500  {object} ErrorResponse
//// @Router       /api/v1/tryout/exam/{id} [delete]
//// DeleteExam removes an exam by its ID.
//func (ctr *controller) DeleteExam(c *fiber.Ctx) error {
//	id := c.Params("id")
//	if err := ctr.service.RemoveExam(id); err != nil {
//		c.Status(http.StatusInternalServerError)
//		return c.JSON(CreateErrorResponse(err))
//	}
//	return c.JSON(id)
//}
