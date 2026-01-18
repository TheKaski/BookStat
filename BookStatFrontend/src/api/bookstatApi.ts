import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { type BookMetadata, type ShelfMetadata } from "../types/domain.types";

export const bookstatApi = createApi({
  reducerPath: "bookstatApi",
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:8080" }),
  endpoints: (builder) => ({
    getAllBooks: builder.query({
      query: () => "/books/",
    }),

    getBookById: builder.query({
      query: (id) => `/books/${id}`,
    }),

    getUnassignedBooks: builder.query<BookMetadata[], void>({
      query: () => "/books/unassigned",
    }),

    createBook: builder.mutation({
      query: (book) => ({
        url: `/books/`,
        method: "POST",
        body: book,
      }),
    }),

    updateBookById: builder.mutation({
      query: (book) => ({
        url: `/books/${book.id}`,
        method: "PUT",
        body: book,
      }),
    }),

    deleteBookById: builder.mutation({
      query: (id) => ({
        url: `/books/${id}`,
        method: "DELETE",
      }),
    }),

    getAllShelves: builder.query<ShelfMetadata[], void>({
      query: () => "/shelves",
    }),

    getShelfById: builder.query({
      query: (id) => `/shelves/${id}`,
    }),

    createShelf: builder.mutation({
      query: (shelf) => ({
        url: `/shelves/`,
        method: "POST",
        body: shelf,
      }),
    }),

    updateShelfById: builder.mutation({
      query: (shelf) => ({
        url: `/shelves/${shelf.id}`,
        method: "PUT",
        body: shelf,
      }),
    }),

    deleteShelfById: builder.mutation({
      query: (id) => ({
        url: `/shelves/${id}`,
        method: "DELETE",
      }),
    }),
  }),
});

export const {
  useGetAllBooksQuery,
  useGetBookByIdQuery,
  useGetUnassignedBooksQuery,
  useCreateBookMutation,
  useUpdateBookByIdMutation,
  useDeleteBookByIdMutation,

  useGetAllShelvesQuery,
  useGetShelfByIdQuery,
  useCreateShelfMutation,
  useUpdateShelfByIdMutation,
  useDeleteShelfByIdMutation,
} = bookstatApi;
