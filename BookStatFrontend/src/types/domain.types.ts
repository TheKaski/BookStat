// =========================
// Book Types
// =========================

export interface Book {
  id: number;
  shelfId: number;
  isbn: string;
  title: string;
  author: string;
  language: string;
  numberOfPages: number;
  format: string;
  coverType: string;
  coverImageUrl: string;
  weight: number;
}

export interface BookMetadata {
  id: number;
  shelfId: number;
  title: string;
  author: string;
  coverImageUrl: string;
}

// =========================
// Shelf Types
// =========================

export interface Shelf {
  id: number;
  name: string;
  books?: BookMetadata[]; // matches `omitempty`
}

export interface ShelfMetadata {
  id: number;
  name: string;
}
