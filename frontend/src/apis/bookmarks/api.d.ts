export declare namespace api.bookmarks.response {
    export interface bookmarks {
        id: string
        inProcess: boolean
        url: string
        icon: string
        title: string
        remark: string
        viewTime: string
        sort: number
        createAt: string
    }
}
declare namespace api.bookmarks.request {
    export interface bookmarks {
        url: string
        title: string
        remark: string
        sort: number
    }
}