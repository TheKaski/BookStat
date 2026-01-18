import { useGetUnassignedBooksQuery } from "../api/bookstatApi";

export default function UnassignedBooksContainer() {
  const { data: books, isLoading, isError } = useGetUnassignedBooksQuery();

  if (isLoading) {
    return <p className="text-gray-400 text-center">Ladataan kirjoja...</p>;
  }

  if (isError) {
    return (
      <p className="text-red-500 text-center">Virhe ladattaessa kirjoja</p>
    );
  }

  const isEmpty = books?.length === 0;

  return (
    <div className="p-4">
      <h2 className="text-xl font-semibold">Lajittelemattomat kirjat</h2>

      {isEmpty && <p>No unnassigned books at the moment</p>}
    </div>
  );
}
