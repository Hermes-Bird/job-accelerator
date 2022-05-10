import {makeAutoObservable} from "mobx";
import {AuthDto, LoginDto, RegisterEmployeeDto, UserType} from "../types";
import axiosInstance from "../axios";
import {AxiosError} from "axios";
import {makePersistable, isHydrated} from "mobx-persist-store";

class AuthStore {
  id: number | null = null
  type: UserType | null = null
  accessToken: string | null = null
  refreshToken: string | null = null
  isPersist = false

  constructor() {
    makeAutoObservable(this)
    makePersistable(this, {
      name: 'AuthStore',
      storage: window.localStorage,
      properties: [
        'id', 'type', 'accessToken', 'refreshToken'
      ],
    }).then(() => {
      this.updateAccessToken()
    })
  }

  get isHydrated(): boolean {
    return isHydrated(this)
  }

  Logout() {
    this.id = null
    this.type = null
    this.refreshToken = null
    this.accessToken = null
  }

  RegisterCompany() {

  }

  async RegisterEmployee(emplDto: RegisterEmployeeDto) {
    try {
      const { data: {
        id, type,
        access_token, refresh_token
      }} = await axiosInstance.post<AuthDto>('/auth/employee/register', emplDto)
      this.accessToken = access_token
      this.refreshToken = refresh_token
      this.type = type
      this.id = id

      this.updateAccessToken()
    } catch (e) {
      console.error(e)
    }

  }

  async LoginUser(userType: UserType, loginDto: LoginDto) {
    try {
      const { data: {
        id, type,
        access_token, refresh_token
      }} = await axiosInstance.post<AuthDto>(`auth/${userType}/login`, loginDto)
      this.accessToken = access_token
      this.refreshToken = refresh_token
      this.type = type
      this.id = id

      this.updateAccessToken()
    } catch (e: any) {
      const err = e as AxiosError
      console.error(err)
    }
  }

  updateAccessToken() {
    axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${this.accessToken}`
  }
}

export default new AuthStore()