import { configureStore } from "@reduxjs/toolkit";
import { bookstatApi } from "./api/bookstatApi";

export const store = configureStore({
  reducer: {
    [bookstatApi.reducerPath]: bookstatApi.reducer,
  },
  middleware: (getDefault) => getDefault().concat(bookstatApi.middleware),
});
