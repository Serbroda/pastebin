import ApiService from "./ApiService";
import {CreatePasteDto} from "../models/CreatePasteDto";
import {UpdatePasteDto} from "../models/UpdatePasteDto";
import {PasteDto} from "../models/PasteDto";

export default class PasteService extends ApiService {

    async getPastes(): Promise<PasteDto[]> {
        const response = await fetch(super.apiUrl(`/v1/pastes`));
        return response.json();
    }

    async createPaste(dto: CreatePasteDto): Promise<PasteDto> {
        const response = await fetch(super.apiUrl(`/v1/pastes`), {
            method: 'post',
            body: JSON.stringify(dto),
            headers: {
                'Content-Type': 'application/json'
            },
        });
        return response.json();
    }

    async getPaste(id: string): Promise<PasteDto> {
        const response = await fetch(super.apiUrl(`/v1/pastes/${id}`));
        return response.json();
    }

    async updatePaste(id: string, dto: UpdatePasteDto): Promise<PasteDto> {
        const response = await fetch(super.apiUrl(`/v1/pastes/${id}`), {
            method: 'patch',
            body: JSON.stringify(dto),
            headers: {
                'Content-Type': 'application/json'
            },
        });
        return response.json();
    }

    async deletePaste(id: string): Promise<void> {
        await fetch(super.apiUrl(`/v1/pastes/${id}`), {
            method: 'delete'
        });
    }
}
