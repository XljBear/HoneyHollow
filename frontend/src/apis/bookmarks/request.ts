import request from '../axios'
import { api } from './api'

export const getBookmarksList = (isView: boolean) =>
    request.get(`/bookmarks${(isView ? "?v=true" : "")}`) as Promise<api.bookmarks.response.bookmarks[]>
export const getBookmarksById = (bookmarksId: string) =>
    request.get(`/bookmarks/${bookmarksId}`) as Promise<api.bookmarks.response.bookmarks>
export const viewBookmarks = (bookmarksId: string) =>
    request.put(`/view/${bookmarksId}`)
export const unViewBookmarks = (bookmarksId: string) =>
    request.put(`/unView/${bookmarksId}`)
export const deleteBookmarks = (bookmarksId: string) =>
    request.delete(`/bookmarks/${bookmarksId}`)
export const getUnViewBookmarksCount = () =>
    request.get('/bookmarks/count') as Promise<number>
export const createBookmarks = (bookmarks: api.bookmarks.request.bookmarks) =>
    request.post('/bookmarks', bookmarks)
export const updateBookmarks = (bookmarks: api.bookmarks.request.bookmarks) =>
    request.put('/bookmarks', bookmarks)