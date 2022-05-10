enum RouteTypes {
  Home = '/',
  Vacancy = '/vacancy/:id',
  Vacancies = '/vacancies',
  FavoriteVacancies = 'employee/favorites',
  CreateVacancy = '/create/vacancy',
  Company = '/company/:id',
  Employee = '/employee/:id',
  AuthEmployee = '/auth/employee',
  AuthCompany = '/auth/company',
  EditEmployee = '/edit/employee',
  EditCompany = '/edit/company',
  EditVacancy = '/edit/vacancy/:id',
  SearchVacancies = '/search/vacancies',
  SearchEmployees = '/search/employees'
}

export default RouteTypes